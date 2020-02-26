package configmap

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

func WatchAndUpload(namespace, directoryName, configMapName string, events chan fsnotify.Event, errors chan error) {
	log.Println("Synchronize", directoryName, "->", configMapName)
	go func() {
		for {
			select {
			case event, ok := <-events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					Upload(namespace, directoryName, configMapName)
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file:", event.Name)
					Upload(namespace, directoryName, configMapName)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("removed file:", event.Name)
					Upload(namespace, directoryName, configMapName)
				}
			case err, ok := <-errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
}

func Upload(namespace, directoryName, configMapName string) {
	configMap := buildConfigMap(namespace, directoryName, configMapName)
	var tcm = &corev1.ConfigMap{}
	if err := client.Get(ctx, *configMap.Metadata.Namespace, *configMap.Metadata.Name, tcm); err != nil {
		log.Println("Create Configmap",*configMap.Metadata.Namespace, *configMap.Metadata.Name)
		if err := client.Create(ctx, configMap); err != nil {
			log.Panicln(err)
		}
	} else {
		log.Println("Update Configmap",*configMap.Metadata.Namespace, *configMap.Metadata.Name)
		if err := client.Update(ctx, configMap); err != nil {
			log.Panicln(err)
		}
	}
}

func buildConfigMap(namespace, directoryName, configMapName string) (cm *corev1.ConfigMap) {
	time.Sleep(5 * time.Second)
	log.Println("Rebuild configmap", configMapName)
	cm = &corev1.ConfigMap{
		Metadata: &metav1.ObjectMeta{
			Name:      &configMapName,
			Namespace: &namespace,
		},
		Data: map[string]string{},
	}
	filepath.Walk(directoryName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if content, err := ioutil.ReadFile(path); err != nil {
			return err
		} else {
			key := strings.Replace(path, directoryName, "", 1)
			if key[0] == '/' {
				key = key[1:]
			}
			cm.Data[key] = string(content)
		}
		return nil
	})
	return cm
}

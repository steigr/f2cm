package configmap

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

func Download(namespace, configMapName, directoryName string) {
	var cm = &corev1.ConfigMap{}

	if err := client.Get(ctx, namespace, configMapName, cm); err != nil {
		log.Panicln(err)
	}

	if _, err := os.Stat(directoryName); err != nil {
		os.MkdirAll(directoryName, 0755)
	}

	for filename := range cm.Data {
		log.Println("Saving",fmt.Sprintf("%s/%s",directoryName,filename),"from",fmt.Sprintf("%s/%s",*cm.Metadata.Namespace,*cm.Metadata.Name))
		if err := ioutil.WriteFile(fmt.Sprintf("%s/%s", directoryName, filename), []byte(cm.Data[filename]), 0644); err != nil {
			log.Panicln(err)
		}
	}
}

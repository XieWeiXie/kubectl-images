package main

import (
	"context"
	"encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ImageEntity struct {
	Namespace       string `json:"namespace"`
	Pod             string `json:"pod"`
	Container       string `json:"container"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
}

const defaultKubeConfig = "/.kube/config"

type K8sConfig struct {
	Namespace string `json:"namespace"`
}

func NewK8sConfig(namespace string) *K8sConfig {
	return &K8sConfig{namespace}
}

func (k *K8sConfig) RestConfig() (*rest.Config, error) {
	home, _ := os.UserHomeDir()
	cfgPath := home + defaultKubeConfig
	return clientcmd.BuildConfigFromFlags("", cfgPath)
}

func (k *K8sConfig) DefaultClient() *kubernetes.Clientset {
	cfg, err := k.RestConfig()
	if err != nil {
		return nil
	}
	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil
	}
	return client
}

func (k *K8sConfig) ListPodImages() []ImageEntity {
	list, err := k.DefaultClient().CoreV1().Pods(k.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil
	}
	var imageEntities []ImageEntity
	for _, one := range list.Items {
		for _, item := range one.Spec.Containers {
			entity := ImageEntity{
				Namespace:       one.Namespace,
				Pod:             one.Name,
				Container:       item.Name,
				Image:           item.Image,
				ImagePullPolicy: string(item.ImagePullPolicy),
			}
			imageEntities = append(imageEntities, entity)
		}
	}
	return imageEntities
}

type PrettyPrint interface {
	Render()
}

func NewPrettyPrint(entities []ImageEntity, strategy string) PrettyPrint {
	switch strategy {
	case "j", "json":
		return JsonPrettyPrint{entities: entities}
	case "y", "yaml":
		return YamlPrettyPrint{entities: entities}
	default:
		return TablePrettyPrint{entities: entities}
	}
}

type JsonPrettyPrint struct {
	entities []ImageEntity
}

func (j JsonPrettyPrint) Render() {
	output, err := json.MarshalIndent(j.entities, "", " ")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to marshal JSON data, err: %v", err)
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(os.Stdout, string(output))
}

type YamlPrettyPrint struct {
	entities []ImageEntity
}

func (y YamlPrettyPrint) Render() {
	output, err := yaml.Marshal(y.entities)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to marsha Yaml data, err: %v", err)
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(os.Stdout, string(output))
}

type TablePrettyPrint struct {
	entities []ImageEntity
}

func (t TablePrettyPrint) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Namespace", "Container", "Pod", "Images", "ImagePullPolicy"})
	table.SetAutoFormatHeaders(false)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	for _, one := range t.entities {
		table.Append([]string{one.Namespace, one.Container, one.Pod, one.Image, one.ImagePullPolicy})
	}
	table.Render()
}

func main() {

	namespace := flag.String("n", "default", "namespace")
	output := flag.String("o", "t", "output: json、yaml、table or short alias j、y、t")
	flag.Parse()

	cfg := NewK8sConfig(*namespace)
	images := cfg.ListPodImages()
	if len(images) > 0 {
		NewPrettyPrint(images, *output).Render()
	}

}

package labels

import (
	"io/ioutil"

	"github.com/pkg/errors"
	apps "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

func getDeployment(path string) (*apps.Deployment, error) {
	yml, err := getYAML(path)
	if err != nil {
		return nil, err
	}
	return decode(yml)
}

func getYAML(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func decode(yaml []byte) (*apps.Deployment, error) {
	scheme := runtime.NewScheme()
	if err := apps.AddToScheme(scheme); err != nil {
		return nil, err
	}
	factory := serializer.NewCodecFactory(scheme)
	decoder := factory.UniversalDeserializer()

	var obj interface{}
	obj, _, err := decoder.Decode(yaml, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "decoding YAML")
	}

	switch obj.(type) {
	case *apps.Deployment:
		deployment := obj.(*apps.Deployment)
		return deployment, nil
	default:
		return nil, errors.New("unknown kind")
	}

	return nil, nil
}

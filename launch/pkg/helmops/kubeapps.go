package kubeapps

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

var KubeappsHost = os.Getenv("KUBEAPPS_SERVER_IP")

type AppRepoResponse struct {
	AppRepo AppRepoItems `json:"appRepository"`
}

type RefreshAppRepoResponse struct {
	AppRepo AppRepository `json:"appRepository"`
}

type AppRepoItems struct {
	Items []AppRepository `json:"items"`
}

type AppRepository struct {
	Metadata AppRepositoryMetadata `json:"metadata"`
	Spec     AppRepositorySpec     `json:"spec"`
}

type AppRepositorySpec struct {
	URL string `json:"url"`
}

type AppRepositoryMetadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

/*
 * Register App Repository -> Look For Robot -> Look For Version
 */

func GetAppRepository(token string, cluster string, namespace string, name string) (AppRepository, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", KubeappsHost+"/api/v1/clusters/"+cluster+"/apprepositories", nil)
	if err != nil {
		return AppRepository{}, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		return AppRepository{}, err
	}
	var appRepositories AppRepoResponse
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AppRepository{}, err
	}

	err = json.Unmarshal(body, &appRepositories)
	if err != nil {
		return AppRepository{}, err
	}

	for _, repo := range appRepositories.AppRepo.Items {
		if repo.Metadata.Name == name {
			return repo, nil
		}
	}
	return AppRepository{}, errors.New("app repository not found")

}

func RefreshAppRepository(token string, cluster string, namespace string, name string) (AppRepository, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", KubeappsHost+"/api/v1/clusters/"+cluster+"/namespaces/"+namespace+"/apprepositories/"+name+"/refresh", nil)
	if err != nil {
		return AppRepository{}, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		return AppRepository{}, err
	}
	var appRepository RefreshAppRepoResponse
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AppRepository{}, err
	}

	err = json.Unmarshal(body, &appRepository)
	if err != nil {
		return AppRepository{}, err
	}

	return appRepository.AppRepo, nil

}

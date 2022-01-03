package kubeapps

import (
	"bytes"
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

type CreateReleaseBody struct {
	AppRepositoryResourceName      string `json:"appRepositoryResourceName"`
	AppRepositoryResourceNamespace string `json:"appRepositoryResourceNamespace"`
	ChartName                      string `json:"chartName"`
	ReleaseName                    string `json:"releaseName"`
	Version                        string `json:"version"`
	Values                         string `json:"values"`
}

type CreateReleaseResponse struct {
	Data    ReleaseInfo `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type ReleaseInfo struct {
	Name string `json:"name"`
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

func CreateRelease(token string, cluster string, namespace string, release CreateReleaseBody) (CreateReleaseResponse, error) {

	client := &http.Client{}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(release)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	req, err := http.NewRequest("POST", KubeappsHost+"/api/kubeops/v1/clusters/"+cluster+"/namespaces/"+namespace+"/releases", &buf)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	var createReleaseResp CreateReleaseResponse
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	err = json.Unmarshal(body, &createReleaseResp)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	if createReleaseResp.Code != 0 {
		// returns 401 if cannot be created
		return CreateReleaseResponse{}, errors.New(createReleaseResp.Message)
	}

	return createReleaseResp, nil

}

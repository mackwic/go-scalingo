package scalingo

import (
	"encoding/json"
	"time"

	"github.com/Scalingo/go-scalingo/http"
	errgo "gopkg.in/errgo.v1"
)

type BackupsService interface {
	BackupList(app string, addonID string) ([]Backup, error)
	BackupDownloadURL(app, addonID, backupID string) (string, error)
}

type BackupStatus string

const (
	BackupStatusScheduled BackupStatus = "scheduled"
	BackupStatusRunning   BackupStatus = "running"
	BackupStatusDone      BackupStatus = "done"
	BackupStatusError     BackupStatus = "error"
)

type Backup struct {
	ID         string       `json:"id"`
	CreatedAt  time.Time    `json:"created_at"`
	Name       string       `json:"name"`
	Size       uint64       `json:"size"`
	Status     BackupStatus `json:"status"`
	DatabaseID string       `json:"database_id"`
	Direct     bool         `json:"direct"`
}

type BackupsRes struct {
	Backups []Backup `json:"backups"`
}

type DownloadURLRes struct {
	DownloadURL string `json:"download_url"`
}

func (c *Client) BackupList(app string, addonID string) ([]Backup, error) {
	var backupRes BackupsRes
	err := c.DBAPI(app, addonID).ResourceList("backups", nil, &backupRes)
	if err != nil {
		return nil, errgo.Notef(err, "fail to get backup")
	}
	return backupRes.Backups, nil
}

func (c *Client) BackupDownloadURL(app, addonID, backupID string) (string, error) {
	req := &http.APIRequest{
		Method: "GET",
		URL:    "/backups/" + backupID + "/archive",
	}
	resp, err := c.DBAPI(app, addonID).Do(req)
	if err != nil {
		return "", errgo.Notef(err, "fail to get backup archive")
	}
	defer resp.Body.Close()

	var downloadRes DownloadURLRes
	err = json.NewDecoder(resp.Body).Decode(&downloadRes)
	if err != nil {
		return "", errgo.Notef(err, "fail to decode backup archive")
	}
	return downloadRes.DownloadURL, nil
}

package HDM

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	fwLocationUrl = "/api/maintenance/fwimage_location"
	firewareTypeUrl = "/api/maintenance/firmware/type"
	flashUrl = "/api/maintenance/flash"
	dwldfwimgUrl = "/api/maintenance/firmware/dwldfwimg"
	progressUrl = "/api/maintenance/firmware/dwldfwstatus-progress"
	verificationUrl = "/api/maintenance/firmware/verification"
	upgradeUrl = "/api/maintenance/firmware/upgrade"
	flashStatusUrl = "/api/maintenance/firmware/flash_progress"
)



type flashSize struct {
	ImaSize int `json:"ima_size"`
}

type HDMType struct {
	PreserveConfig int `json:"preserve_config"`
	FwTypeName string `json:"fw_type_name"`
	RebootType int `json:"reboot_type"`
	RebootTime int `json:"reboot_time"`
}

type localtionJson struct {
	ProtocolType string `json:"protocol_type"`
	ServerAddress string `json:"server_address"`
	ImageName string `json:"image_name"`
	RetryCount int `json:"retry_count"`
}

type flashStatus struct {
	FlashStatus uint8 `json:"flash_status"`
}

type processStatus struct {
	//ID       int    `json:"id"`
	//Action   string `json:"action"`
	//Progress string `json:"progress"`
	State    uint8    `json:"state"`
}


type flashStatusJson struct {
	//Action   string `json:"action"`
	Progress string `json:"progress"`
	//Cc       int    `json:"cc"`
}

func (h *HDM)reqFwimageLocation(filename,tftpIP string) error {

	data := localtionJson {
		ImageName: filename,
		ProtocolType: "tftp",
		RetryCount: 2,
		ServerAddress: tftpIP,
	}
	payloadBytes, _ := json.Marshal(data)

	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("PUT", h.url(fwLocationUrl), body)

	req.Header.Set("Origin", h.baseUrl)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Csrftoken",  h.cSRFToken)

	resp, err := h.http.Do(req)
	if err != nil {
		return errors.New("err: while do the firmware request")
	}
	defer resp.Body.Close()
	respBody,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	return nil
}

func (h *HDM)reqType() error {

	data := HDMType{
		PreserveConfig: 0,
		FwTypeName: "HDM",
		RebootTime: 0,
		RebootType: 0,
	}
	payloadBytes, _ := json.Marshal(data)

	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("POST", h.url(firewareTypeUrl), body)

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("Origin",  h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("X-Csrftoken",  h.cSRFToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	return nil
}



func (h *HDM) reqFlash() error {
	data := flashSize{
		ImaSize: 0,
	}
	payloadBytes, _ := json.Marshal(data)

	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("PUT", h.url(flashUrl), body)

	req.Header.Set("Origin", h.baseUrl)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Csrftoken",  h.cSRFToken)

	resp, err := h.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	return nil
}


func (h *HDM)reqDwldfwimg() error {
	type Payload struct {
		PROTOTYPE int `json:"PROTO_TYPE"`
	}

	data := Payload {
		PROTOTYPE: 1,
	}
	payloadBytes, _ := json.Marshal(data)

	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("PUT", h.url(dwldfwimgUrl), body)

	req.Header.Set("Origin", h.baseUrl)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Csrftoken",  h.cSRFToken)

	resp, err := h.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	return nil
}


func (h *HDM)reqProcess() error {

	req, _ := http.NewRequest("GET", h.url(progressUrl), nil)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Csrftoken",  h.cSRFToken)


	var data processStatus

	for {
		resp, err := h.http.Do(req)
		if err != nil {
			return err
		}

		respBody,_ := ioutil.ReadAll(resp.Body)


		if err := json.Unmarshal(respBody, &data); err != nil {
			return err
		}

		if data.State != 1 && data.State != 2 {
			return errors.New(string(respBody))
		} else if data.State == 2 {
			break
		} else {
			fmt.Println(string(respBody))
		}
		_ = resp.Body.Close()
		time.Sleep(3*time.Second)
	}
	return nil

}

func (h *HDM)reqVerification() error {
	req, _ := http.NewRequest("GET", h.url(verificationUrl), nil)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Csrftoken",  h.cSRFToken)

	resp, err := h.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

	return nil
}

func (h *HDM) Up(filename,tftpIP string) error {

	if err := h.reqFwimageLocation(filename, tftpIP);err != nil {
		return errors.New(fmt.Sprintf("fireware location|err :%s",err.Error()))
	}
	if err := h.reqType();err != nil {
		return errors.New(fmt.Sprintf("update type|err :%s",err.Error()))
	}
	if err := h.reqFlash();err != nil {
		return errors.New(fmt.Sprintf("flash size send|err :%s",err.Error()))
	}
	if err := h.reqDwldfwimg();err != nil {
		return errors.New(fmt.Sprintf("download bin|err :%s",err.Error()))
	}

	if err := h.reqProcess();err != nil {
		return err
	}

	time.Sleep(2*time.Second)
	if err := h.reqVerification();err != nil {
		return err
	}

	if err := h.reqUpgrade();err != nil {
		return errors.New(fmt.Sprintf("upgrade|err :%s",err.Error()))
	}


	if err := h.flashStatus();err != nil {
		return errors.New(fmt.Sprintf("flash write|err :%s",err.Error()))
	}

	return nil
}


func (h *HDM) reqUpgrade() error{

	data := flashStatus {
		FlashStatus: 1,
	}
	payloadBytes, _ := json.Marshal(data)

	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("PUT", h.url(upgradeUrl), body)

	req.Header.Set("Origin", h.baseUrl)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Csrftoken", h.cSRFToken)

	resp, err := h.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	return nil
}


func (h *HDM)flashStatus() error {
	req, _ := http.NewRequest("GET", h.url(flashStatusUrl), nil)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Csrftoken",  h.cSRFToken)

	var data flashStatusJson

	for {
		resp, err := h.http.Do(req)
		if err != nil {
			return err
		}

		respBody,_ := ioutil.ReadAll(resp.Body)


		if err := json.Unmarshal(respBody, &data); err != nil {
			return err
		}

		if data.Progress != "" && strings.Contains(data.Progress, "Complete") {
			break
		} else {
			fmt.Println(string(respBody))
		}
		_ = resp.Body.Close()
		time.Sleep(2*time.Second)
	}

	return nil
}


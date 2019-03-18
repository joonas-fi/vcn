/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package main

import (
	"fmt"
)

type ArtifactTrackingEventRequest struct {
	Client   string `json:"client"`
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
	Url      string `json:"url"`
}

type PublisherTrackingEventRequest struct {
	Name string `json:"name"`
}

func TrackVerify(hash string, filename string) (err error) {
	restError := new(Error)
	token, err := LoadToken()
	if err != nil {
		return err
	}
	r, err := NewSling(token).
		Post(TrackingEvent() + "/verify").
		BodyJSON(ArtifactTrackingEventRequest{
			Client:   VcnClientName(),
			Filename: filename,
			Hash:     hash,
		}).Receive(nil, restError)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("TrackVerify failed: %s", restError)
	}
	return nil
}

func TrackPublisher(event string) (err error) {
	restError := new(Error)
	token, err := LoadToken()
	if err != nil {
		return err
	}
	r, err := NewSling(token).
		Post(TrackingEvent() + "/publisher").
		BodyJSON(PublisherTrackingEventRequest{
			Name: event,
		}).Receive(nil, restError)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("TrackPublisher failed: %s", restError)
	}
	return nil
}
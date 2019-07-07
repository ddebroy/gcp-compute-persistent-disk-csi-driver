/*
Copyright 2019 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"encoding/json"
	"net"
)

func NewReader(c net.Conn) *Reader {
	return &Reader{
		d: json.NewDecoder(c),
	}
}

type Reader struct {
	d *json.Decoder
}

func (r *Reader) ReadHeader() (*Message, error) {
	var msg Message
	if err := r.d.Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
func (r *Reader) ReadWithBody(body interface{}) (*Message, error) {
	msg, err := r.ReadHeader()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return nil, err
	}
	return msg, nil
}

func NewWriter(c net.Conn) *Writer {
	return &Writer{
		e: json.NewEncoder(c),
	}
}

type Writer struct {
	e *json.Encoder
}

func (r *Writer) Write(op Operation, v interface{}) error {
	m, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return r.e.Encode(Message{
		Type: op,
		Body: m,
	})
}

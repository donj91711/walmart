package main

import (
	"reflect"
	"testing"
)

func Test_getConfig(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		{
			name: "test getConfig",
			want: Config{
				RingSize: 100,
				FileName: "data.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRingBuffer(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *RingBuffer
	}{
		{
			name: "test NewRingBuffer",
			args: args{size: 1},
			want: &RingBuffer{
				buffer: []Data{
					{},
				},
				size:     1,
				readIdx:  0,
				writeIdx: 0,
				count:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRingBuffer(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRingBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRingBuffer_Push(t *testing.T) {
	type fields struct {
		buffer   []Data
		size     int
		readIdx  int
		writeIdx int
		count    int
	}
	type args struct {
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test Push",
			fields: fields{
				buffer: []Data{
					{},
				},
				size:     1,
				readIdx:  0,
				writeIdx: 0,
				count:    0,
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RingBuffer{
				buffer:   tt.fields.buffer,
				size:     tt.fields.size,
				readIdx:  tt.fields.readIdx,
				writeIdx: tt.fields.writeIdx,
				count:    tt.fields.count,
			}
			rb.Push(tt.args.data)
		})
	}
}

func TestRingBuffer_Pop(t *testing.T) {
	type fields struct {
		buffer   []Data
		size     int
		readIdx  int
		writeIdx int
		count    int
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
		want1  bool
	}{
		{
			name: "test Pop",
			fields: fields{
				buffer: []Data{
					{},
				},
				size:     1,
				readIdx:  0,
				writeIdx: 0,
				count:    0,
			},
			want:  Data{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RingBuffer{
				buffer:   tt.fields.buffer,
				size:     tt.fields.size,
				readIdx:  tt.fields.readIdx,
				writeIdx: tt.fields.writeIdx,
				count:    tt.fields.count,
			}
			got, got1 := rb.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RingBuffer.Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RingBuffer.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

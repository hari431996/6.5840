package kvsrv

import (
	"log"
	"sync"
)

const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}


type KVServer struct {
	mu sync.Mutex
	DataStore map[string]string

	// Your definitions here.
}


func (kv *KVServer) Get(args *GetArgs, reply *GetReply)  {
	// Your code here.
	key := args.Key
	reply.Value = ""
	
	kv.mu.Lock()
	if val, ok := kv.DataStore[key]; ok{
		reply.Value = val
		
	}
	kv.mu.Unlock()

	
}

func (kv *KVServer) Put(args *PutAppendArgs, reply *PutAppendReply)  {
	// Your code here.


	key := args.Key
	val := args.Value
	
	kv.mu.Lock()
	kv.DataStore[key] = val
	kv.mu.Unlock()

	reply.Value = "ok!!!"
	
	
}

func (kv *KVServer) Append(args *PutAppendArgs, reply *PutAppendReply)  {
	// Your code here.
	key := args.Key
	value := args.Value

	kv.mu.Lock()
	if val, ok := kv.DataStore[key];ok{
		kv.DataStore[key]+=value
		reply.Value = val
		

	}else{
		kv.DataStore[key] = value
		reply.Value = value
	}
	kv.mu.Unlock()

	
	
	

}

func StartKVServer() *KVServer {
	
	kv := &KVServer{}
	

	// You may need initialization code here.
	kv.DataStore = make(map[string]string)
	kv.mu = sync.Mutex{}

	return kv
}

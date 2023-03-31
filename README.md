# INSTRUCTIONS
## CLONE LLAMA-SAAS
besides having llama.cpp, you will need to clone https://github.com/avilum/llama-saas somewhere else in your system  
## AT llama-saas/server.go  
change LLAMA_MAIN (to the main in llama.cpp), LLAMA_MODEL_PATH, LLAMA_NUM_THREADS (to 4) and MAX_OUTPUT_LENGTH approprietly  
## AT tekbot/
modify LLAMA_SAAS_PATH in ./answer.sh and ./run.sh  
you may need to chmod +x ./answer.sh and ./llama.py
## AT tekbot/examples/commands/main.go
give it a unique irc_nick and a bot_name(which will trigger it)  
## AT tekbot/examples/commands/dev.toml
give it a unique irc_nick  
# RUN IT
well, you could run ./run.sh, if you do just close the terminal from where you did  
however!  
I rather open two terminals and on each I run:
```
cd tekbot
cd ./examples/commands/
go run main.go -config dev.toml 
```
and
```
cd llama-saas
go build
./server 
```



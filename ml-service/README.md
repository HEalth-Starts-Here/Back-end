Setup steps:
---
~~~bash
cd SERVICE_NAME
python -m venv .venv
source .venv/bin/activate
python -m pip install -U pip
pip install -r requirements.txt
~~~

For text summarization:
~~~bash
python -m spacy download ru_core_news_sm
~~~

Text diarization also requires the command-line tool `ffmpeg` to be installed on your system:
~~~bash
sudo apt update && sudo apt install ffmpeg
~~~

Run service:
---
~~~bash
python main.py configs/iqa_config.yml
~~~

Generate grpc files:
---
NOTE: данную команду необходимо запускать из папки api соответствующего микросервиса
~~~bash
python -m grpc_tools.protoc -I../../protos --python_out=. --pyi_out=. --grpc_python_out=. ../../protos/affected_area.proto
~~~

Setup steps:
~~~bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
~~~

Run service:
~~~bash
python main.py configs/iqa_config.yml
~~~

Generate grpc files:
~~~bash
python -m grpc_tools.protoc -I./api --python_out=./protos --pyi_out=./protos --grpc_python_out=./protos ./api/iqa.proto
~~~

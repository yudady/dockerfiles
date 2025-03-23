# 容器化部署


ollama-deep-researcher

API: http://127.0.0.1:2024

Docs: http://127.0.0.1:2024/docs

LangGraph Studio Web UI: https://smith.langchain.com/studio/?baseUrl=http://127.0.0.1:2024


```shell


2025-03-17T04:00:16.018614726Z    Building ollama-deep-researcher @ file:///app
2025-03-17T04:00:16.223945692Z Downloading sqlalchemy (3.1MiB)
2025-03-17T04:00:16.224285985Z Downloading lxml (4.6MiB)
2025-03-17T04:00:16.224867319Z Downloading numpy (13.7MiB)
2025-03-17T04:00:16.225818363Z Downloading langchain-community (2.4MiB)
2025-03-17T04:00:16.225851696Z Downloading zstandard (4.7MiB)
2025-03-17T04:00:16.226299697Z Downloading pydantic-core (1.7MiB)
2025-03-17T04:00:16.226396364Z Downloading tiktoken (1.1MiB)
2025-03-17T04:00:16.226628781Z Downloading cryptography (3.6MiB)
2025-03-17T04:00:16.228213242Z Downloading aiohttp (1.6MiB)
2025-03-17T04:00:16.355853022Z Downloading jsonschema-rs (1.9MiB)
2025-03-17T04:00:16.558195816Z    Building primp==0.14.0
2025-03-17T04:00:16.754138057Z  Downloaded tiktoken
2025-03-17T04:00:16.911575268Z  Downloaded aiohttp
2025-03-17T04:00:16.919601075Z  Downloaded pydantic-core
2025-03-17T04:00:16.950925091Z  Downloaded jsonschema-rs
2025-03-17T04:00:17.015772004Z  Downloaded langchain-community
2025-03-17T04:00:17.131654220Z  Downloaded sqlalchemy
2025-03-17T04:00:17.199263930Z  Downloaded cryptography
2025-03-17T04:00:17.307496215Z  Downloaded lxml
2025-03-17T04:00:17.311399806Z  Downloaded zstandard
2025-03-17T04:00:17.381424020Z       Built ollama-deep-researcher @ file:///app
2025-03-17T04:00:17.569838164Z   × Failed to build `primp==0.14.0`
2025-03-17T04:00:17.569850789Z   ├─▶ The build backend returned an error
2025-03-17T04:00:17.569852080Z   ╰─▶ Call to `maturin.build_wheel` failed (exit status: 1)
2025-03-17T04:00:17.569852955Z 
2025-03-17T04:00:17.569853664Z       [stdout]
2025-03-17T04:00:17.569854372Z       Running `maturin pep517 build-wheel -i
2025-03-17T04:00:17.569855122Z       /root/.cache/uv/builds-v0/.tmp3XqIpL/bin/python --compatibility off`
2025-03-17T04:00:17.569855789Z 
2025-03-17T04:00:17.569856414Z       [stderr]
2025-03-17T04:00:17.569857122Z       💥 maturin failed
2025-03-17T04:00:17.569857997Z         Caused by: Cargo metadata failed. Do you have cargo in your PATH?
2025-03-17T04:00:17.569858955Z         Caused by: No such file or directory (os error 2)
2025-03-17T04:00:17.569859872Z       Error: command ['maturin', 'pep517', 'build-wheel', '-i',
2025-03-17T04:00:17.569860872Z       '/root/.cache/uv/builds-v0/.tmp3XqIpL/bin/python', '--compatibility',
2025-03-17T04:00:17.569861914Z       'off'] returned non-zero exit status 1
2025-03-17T04:00:17.569878955Z 
2025-03-17T04:00:17.569879914Z       hint: This usually indicates a problem with the package or the build
2025-03-17T04:00:17.569906372Z       environment.
2025-03-17T04:00:17.569907497Z   help: `primp` (v0.14.0) was included because `ollama-deep-researcher`
2025-03-17T04:00:17.569908330Z         (v0.0.1) depends on `duckduckgo-search` (v7.5.2) which depends on
2025-03-17T04:00:17.569909122Z         `primp`


```
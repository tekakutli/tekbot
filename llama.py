import requests
import sys
import subprocess

def main():
    session = requests.Session()
    question=' '.join(sys.argv[1:])
    query=question

    res = session.get(f"http://localhost:8080?text={query}", stream=True)
    output=""
    try:
        for line in res.iter_content(3):
            try:
                x = line.decode()
                if x:
                    print(x, end="", flush=True)
                    output=output+x
                    if  output.count("Anon:")>1:
                        break
                    if  output.count("Miku:")>1:
                        break
            except KeyboardInterrupt:
                break
            except:
                continue
    except:
        ...
    finally:
        print("")


if __name__ == "__main__":
    main()

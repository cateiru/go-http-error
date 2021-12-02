import requests
from lxml.html import fromstring


def main():
    client_status: list[int] = list(range(400, 418 + 1)) + [422, 425, 426, 428, 429, 431, 451]
    server_status: list[int] = list(range(500, 508 + 1)) + [510, 511]

    client_error_go = "package httperror\n\nimport \"net/http\"\n"

    for status in client_status:
        client_error_go += create_go_script(status)
        client_error_go += "\n\n"

    client_error_go += "\n"

    with open("httperror/client_error.go", mode="w") as f:
        f.write(client_error_go)

    client_error_go = "package httperror\n\nimport \"net/http\"\n"

    for status in server_status:
        client_error_go += create_go_script(status)
        client_error_go += "\n\n"

    client_error_go += "\n"

    with open("httperror/server_error.go", mode="w") as f:
        f.write(client_error_go)


def create_go_script(status: int):
    link: str = f"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/{status}"
    res = requests.get(link)
    tree = fromstring(res.content)

    title: str = tree.findtext('.//title').split("-")[0]

    print(title)

    status_name = "".join(title.split(" ")[1:])

    go_text = f"""
// {title}
//
// See more: {link}
func New{status_name}Error(err error) *HTTPError {{
\treturn NewError(http.Status{status_name}, err)
}}"""
    return go_text


if __name__ == "__main__":
    main()

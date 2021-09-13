import glob
import json
import requests
import sys

from bs4 import BeautifulSoup


def update_titles(filename):
    try:
        with open(filename, 'r') as fr:
            data = json.load(fr)

        references = data.get('custom', {}).get('references', {})
        if len(references) == 0:
            return
        for ref in references:
            url = ref['url']
            resp = requests.get(url)
            if resp.status_code != 200:
                print(f'an error occurred processing {url}')

            html = BeautifulSoup(resp.text, features="html.parser")
            title = html.title.text.split('|')
            title = title[0].split(' - ')[0].rstrip()
            ref['title'] = title

        with open(filename, 'w') as fw:
            json.dump(data, fw, indent=4)
            fw.flush()

    except:
        e = sys.exc_info()[0]
        print(f'issue processing {filename}: {e}')


def run():
    for f in glob.glob(f'./*/*/*/metadata.json'):
        update_titles(f)


if __name__ == '__main__':
    run()

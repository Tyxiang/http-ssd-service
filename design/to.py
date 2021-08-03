import xmltodict
import json 

f = open("readme.opml", 'rb')
l = f.read()
d = xmltodict.parse(l)
f.close()

j = json.dumps(d, indent = 4)
f = open('readme.json', 'w')
f.write(j)
f.close()
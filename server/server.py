#-*- coding:utf-8 -*-
from flask import Flask,render_template,request
import MySQLdb,sys
import base64

reload(sys)
sys.setdefaultencoding("utf-8")
MYSQL_HOST = '127.0.0.1'
MYSQL_USER = 'root'
MYSQL_PASS = 'root'
MYSQL_DB   = 'bot'

app = Flask(__name__)
@app.route('/view')
def view():
    conn = MySQLdb.connect(MYSQL_HOST, MYSQL_USER,MYSQL_PASS,MYSQL_DB,charset="utf8")
    cursor = conn.cursor()
    cursor.execute('select * from bot')
    results=cursor.fetchall()
    conn.close()
    return render_template("show.html",
        results = results)
@app.route('/')
def main():
    return 'Hello,World'
@app.route("/callback",methods=[ 'GET'])
def callback():
    conn = MySQLdb.connect(MYSQL_HOST, MYSQL_USER,MYSQL_PASS,MYSQL_DB,charset="utf8")
    cursor = conn.cursor()
    mac = base64.b64decode(request.values['mac'])
    ip  = base64.b64decode(request.values['ip'])
    time = base64.b64decode(request.values['time'])
    print mac,ip,time
    cursor.execute('insert into bot(mac,ip,time) values ("%s","%s","%s")' %(mac,ip,time))
    cursor.close()
    conn.close()
    return 'ok'

if __name__ == '__main__':
    try:
        app.run(host='0.0.0.0',port=80,debug=True,threaded=True)
    except:
        pass
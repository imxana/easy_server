from flask import Flask, url_for, Blueprint

app = Flask(__name__, static_url_path='/..')

# route
@app.route('/')
def index():
    return app.send_static_file('a.html')

@app.route('/img')
def image():
    return app.send_static_file('atago.jpg')

if __name__=='__main__':
    app.run(host='0.0.0.0')

from flask import Flask, render_template
import data

app = Flask(__name__)

@app.route('/')
def hello():
    return render_template("home.html", data=data)
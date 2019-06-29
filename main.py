from flask import Flask, render_template
from datetime import datetime
import data

app = Flask(__name__)

@app.route('/')
def hello():
    return render_template("home.html", data=data, now=datetime.utcnow())
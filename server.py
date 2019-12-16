from flask import Flask
app = Flask(__name__)

@app.route("/")
def index():
    return "You've reached the Pi server index page"


if __name__ == "__main__":
    app.run()
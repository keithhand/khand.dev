FROM python:3.7.3
ADD . /app
WORKDIR /app
RUN pip install flask gunicorn
CMD ["gunicorn", "-b", "0.0.0.0:8000", "main:app"]
EXPOSE 8000
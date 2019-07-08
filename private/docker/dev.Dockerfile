FROM python:3.7.3
ADD . /app
WORKDIR /app
RUN pip install -r requirements.txt
ENTRYPOINT ["flask"]
CMD ["run"]
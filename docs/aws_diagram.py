from diagrams import Diagram
from diagrams.aws.network import CF
from diagrams.aws.network import APIGateway
from diagrams.aws.compute import Lambda
from diagrams.aws.database import RDS

with Diagram("fiber-aws-serverless", filename="./docs/aws_diagram",  show=False):
    CF("for CustomDomain") >> APIGateway("http-gateway") >> Lambda("fiber-app") >> RDS("aws RDS")
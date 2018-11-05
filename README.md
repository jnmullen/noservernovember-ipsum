
# No Server November Challenge Number 1

See this page [NoServerNovember](https://serverless.com/blog/no-server-november-challenge/)

First challenge : 

>Build a Serverless Ipsum generator.
> Build a simple serverless-backed web app that displays Serverless Ipsum when it is loaded. Or Tony Danza Ipsum. Or The Office Ipsum. Or Reasons-I-Can’t-Take-Out-The-Trash Ipsum.
> As long as it looks like Lorem Ipsum, but uses different words, we’re good. The page doesn’t have to look fancy, and you can do this even if you’ve never coded anything in your life!

### Prerequisites

To try out this example you will need to install the Serverless Framework for your platform from here : 

[Install Serverless Framework](https://serverless.com/framework/docs/providers/aws/guide/quick-start/)

Assume you have created a file with your AWS credentials in __~./aws/credentials__ as the Serverless Framework will require these to be able to run.

### Running Serverless Framework

To create the infrastructure run:
```
sls deploy --stage development -v
```
The serverless config has been setup to allow multiple stages/environments to be created but as yet only one called development exists.

Once you have finished you can destroy the infrastructure with : 
```
sls remove --stage development
```

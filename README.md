# auto-blog-wordpress

Automated submission to WordPress using GPT4.

# setup

1. Clone this repository https://github.com/tkwest3143/auto-blog-wordpress.git of git@github.com:tkwest3143/auto-blog-wordpress.git
2. Copy the .env.sample file and rename the file to .env
3. Change the settings in the .env file as follows
   | key | value |
   | ------------- | ------------- |
   | OPENAI_API_KEY | API key created with "https://platform.openai.com/account/api-keys |
   | WORDPRESS_SITE_NAME | URL of the wordpress site to which you are submitting |
   | WORDPRESS_USER_NAME | Site Username |
   | WORDPRESS_API_KEY | Application password created from the site's profile screen |

# start

Execute the following command from a terminal

```
$ make run
```

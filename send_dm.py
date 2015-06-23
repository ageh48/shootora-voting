
import argparse
import commands
import tweepy

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('consumer_key')
    parser.add_argument('consumer_secret')
    parser.add_argument('access_token')
    parser.add_argument('access_token_secret')
    parser.add_argument('target_screen_name')
    parser.add_argument('target_name')
    args = parser.parse_args()

    auth = tweepy.OAuthHandler(args.consumer_key, args.consumer_secret)
    auth.set_access_token(args.access_token, args.access_token_secret)
    api = tweepy.API(auth)

    message = commands.getoutput('./adduser.sh %s' % args.target_name)
    api.send_direct_message(screen_name=args.target_screen_name, text=message)
    print 'To @%s: %s' % (args.target_screen_name, message)

if __name__ == '__main__':
    main()

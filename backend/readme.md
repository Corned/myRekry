# Generating a Gmail App Password

App passwords are special one-time passwords Google lets you generate for scripts and apps that connect to your account directly (e.g. over IMAP), rather than through a browser login.

## Prerequisites

You need **2-Step Verification** enabled on your Google account. If you haven't done this yet:

1. Go to [myaccount.google.com](https://myaccount.google.com)
2. Click **Security** in the left sidebar
3. Under "How you sign in to Google", click **2-Step Verification** and follow the steps

## Generating the password

1. Go to [myaccount.google.com](https://myaccount.google.com)
2. Click **Security** in the left sidebar
3. Under "How you sign in to Google", click **2-Step Verification**
4. Scroll to the bottom and click **App passwords**
5. Enter a name for the password (e.g. `go email script`) and click **Create**
6. Google will display a 16-character password — **copy it immediately**, you won't be able to see it again

## Using it

Store the app password in a `.env` file (never hardcode it in your source code):

```
GMAIL_EMAIL=you@gmail.com
GMAIL_PASSWORD=abcd efgh ijkl mnop
```

The spaces in the password are optional — both formats work.

## Revoking a password

If you think an app password has been leaked, you can revoke it without changing your main Google password:

1. Go back to **Security → 2-Step Verification → App passwords**
2. Click the delete icon next to the password you want to revoke

## Troubleshooting

**I don't see the "App passwords" option**
- Make sure 2-Step Verification is fully enabled
- If your account is managed by a company or school, an admin may have disabled this feature — contact them for access

**Authentication fails even with the app password**
- Make sure IMAP is enabled in Gmail: **Settings → See all settings → Forwarding and POP/IMAP → Enable IMAP**
- Double-check there are no typos — copy-paste the password directly from your `.env` file
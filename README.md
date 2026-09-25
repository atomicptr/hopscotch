# hopscotch

Simple application for redirecting multiple urls from one host to another, configured completely via env vars

## Configuration

Please be sure to pick a unique `{NAME}` for every entry

- `HOPSCOTCH_REDIRECT_{NAME}_FROM=example.com` - Redirect **FROM** example.com
- `HOPSCOTCH_REDIRECT_{NAME}_TO=atomicptr.dev` - Redirect **TO** atomicptr.dev
- `HOPSCOTCH_REDIRECT_{NAME}_PERMANENT=true` - (Optional) Make the redirect permanent (default: false)

For example lets say we have **a.com**, **b.com** and **c.com** and want to redirect them to **atomicptr.dev**

```bash
# a.com -> atomicptr.dev
HOPSCOTCH_REDIRECT_A_COM_FROM=a.com
HOPSCOTCH_REDIRECT_A_COM_TO=atomicptr.dev

# b.com -> atomicptr.dev
HOPSCOTCH_REDIRECT_B_COM_FROM=b.com
HOPSCOTCH_REDIRECT_B_COM_TO=atomicptr.dev

# c.com -> atomicptr.dev
HOPSCOTCH_REDIRECT_C_COM_FROM=c.com
HOPSCOTCH_REDIRECT_C_COM_TO=atomicptr.dev
```

## License

GPLv3

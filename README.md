# KalshiGo

KalshiGo is a Go package providing low-level bindings to the Kalshi API. Inspired by [DiscordGo](https://github.com/bwmarrin/discordgo/tree/master). Add `kalshipulse` on Discord for support or discussions.

## Usage

```golang
c, err := NewFromKeyPath("./kalshi_private_key.pem", "7627b510-20db-4b17-8b8c-83f18c9344a1", "https://api.elections.kalshi.com")

if err != nil {
    panic(err)
}

s, err := c.GetSeries(&GetSeriesParams{
    SeriesTicker: "KXPAYROLLS",
})

if err != nil {
    panic(err)
}

fmt.Println(s)
```

In general, any GET endpoint (series, market, etc) will have a corresponding Get{type}Params struct that needs to be filled.

## Status
| Section      | Support Status |
| ------------ | -------------- |
| market       | ✅             |
| exchange     | ✅             |
| auth [1]     | ✅             |
| collection   | ❌             |
| portfolio    | ❌             |
| sockets      | ❌             |

[1] The auth endpoint has been deprecated and replaced with API key-based authentication, which we support.

# TODO:
- Add features (first priority)

- Add context support

- Add godoc comments

- Add cleaner logging

- Deduplicate code
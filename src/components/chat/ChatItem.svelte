<script>
    import { Tile } from "carbon-components-svelte";

    export let avatar = ""
    export let isAnnouncement = false
    export let announcementColor = ""
    export let username = ""
    export let usernameColor = ""
    export let badges = {
        twitchStaff: false,
        admin: false,
        broadcaster: false,
        moderator: false,
        verified: false,
        vip: false,
        artist: false,
        gameDeveloper: false,
        withoutVideo: false,
        withoutAudio: false,
        turbo: false,
        prime: false,
        bitsBadge: {
            1: false,
            100: false,
            1_000: false,
            5_000: false,
            10_000: false,
            25_000: false,
            50_000: false,
            75_000: false,
            100_000: false,
            200_000: false,
            300_000: false,
            400_000: false,
            500_000: false,
            600_000: false,
            700_000: false,
            800_000: false,
            900_000: false,
            1_000_000: false
        },
        subGift: {
            1: false,
            5: false,
            10: false,
            25: false,
            50: false,
            100: false,
            150: false,
            200: false,
            250: false,
            300: false,
            350: false,
            400: false,
            450: false,
            500: false,
            550: false,
            600: false,
            650: false,
            700: false,
            750: false,
            800: false,
            850: false,
            900: false,
            950: false,
            1_000: false,
            2_000: false,
            3_000: false,
            4_000: false,
            5_000: false
        },
        giftLeaderboard: {
            1: false,
            2: false,
            3: false
        },
        bitLeaderboard: {
            1: false,
            2: false,
            3: false
        }
    }
</script>

<Tile>Content</Tile>

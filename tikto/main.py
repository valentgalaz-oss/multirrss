from TikTokApi import TikTokApi
import os
import asyncio

ms_token = os.environ.get(
    "ms_token", None
) 

async def get_user_data():
    async with TikTokApi() as api:
        await api.create_sessions(
            ms_tokens=[ms_token],
            num_sessions=1,
            sleep_after=3,
            browser=os.getenv("TIKTOK_BROWSER", "chromium"),
        )
        user = api.user("kyliejenner")
        user_data = await user.info()
        print(user_data)

        async for video in user.videos(count=30):
            print(video)
            print(video.as_dict)

        async for playlist in user.playlists():
            print(playlist)

    asyncio.run(get_user_data())
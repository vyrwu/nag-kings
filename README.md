# The Kings 👑

In thesprit of NAG, I will create a nice API for kings data. API should be documented first in Swagger. When application is done, it should be registered in Docker registry and deployed to Azure.   

- /kings/stats: returns analytics on data. Include caching or memoization (not to process too much).
{
	totalKings: number,
	longestRulingKing:  Object{...king},
	longestRulingHouse: string,
	mostCommonFirstName: string
}
- /kings: It will return unchanged kings data. Simple redirect. See `kingsData.json` for example payload.


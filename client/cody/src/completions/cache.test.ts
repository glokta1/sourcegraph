import { CompletionsCache } from './cache'

describe('CompletionsCache', () => {
    it('returns the cached completion items', () => {
        const cache = new CompletionsCache()
        cache.add([{ prefix: 'foo\n', content: 'bar', prompt: '' }])

        expect(cache.get('foo\n')).toEqual([{ prefix: 'foo\n', content: 'bar', prompt: '' }])
    })

    it('returns the cached items when the prefix includes characters from the completion', () => {
        const cache = new CompletionsCache()
        cache.add([{ prefix: 'foo\n', content: 'bar', prompt: '' }])

        expect(cache.get('foo\nb')).toEqual([{ prefix: 'foo\nb', content: 'ar', prompt: '' }])
        expect(cache.get('foo\nba')).toEqual([{ prefix: 'foo\nba', content: 'r', prompt: '' }])
    })

    it('returns the cached items when the prefix has less whitespace', () => {
        const cache = new CompletionsCache()
        cache.add([{ prefix: 'foo \n  ', content: 'bar', prompt: '' }])

        expect(cache.get('foo \n  ')).toEqual([{ prefix: 'foo \n  ', content: 'bar', prompt: '' }])
        expect(cache.get('foo \n ')).toEqual([{ prefix: 'foo \n ', content: 'bar', prompt: '' }])
        expect(cache.get('foo \n')).toEqual([{ prefix: 'foo \n', content: 'bar', prompt: '' }])
        expect(cache.get('foo ')).toEqual(undefined)
    })
})

const createPlugin = require('@extism/extism')

;(async() => {
    const p = await createPlugin('plugin.wasm', {
        useWasi: true,
        enableWasiOutput: true,
    })

    console.log(await p.call("hello", "ppp"))
})()
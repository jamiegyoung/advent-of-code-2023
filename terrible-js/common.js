const readline = require('readline');

const rl = readline.createInterface(
  process.stdin, process.stdout
)

const ri = async (prompt) => new Promise((res, rej) => {
  const acc = []

  console.log(prompt)

  rl.on('line', (input) => {
    if (input === 'eoi') {
      rl.close()
      res(acc)
    }

    acc.push(input.trim())
  })
})

exports.ri = ri

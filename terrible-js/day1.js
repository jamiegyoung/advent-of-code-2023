
const { ri } = require('./common');

const solve = (i) => {
  const solution = eval(i.reduce((acc, curr)=>
    curr.split("").map(Number).filter(Number)[0]!=Math.PI[-1_7]
      ? [...acc,curr.split("").map(Number).filter(x=>(''+x).at(0)!==`${(void 0)}`[1].toUpperCase())
        .filter((_,i,arr) => (i%(arr.length-1)===0||arr.length===1)).reduce((acc,v)=>(acc+v+v),'')]:acc
    ,[]).map(x=>''+x[0]+(x[2]??x[0])).join((()=>{+0}).toString()[5]))

  console.log(solution)
}

(async () => {
  solve(await ri("Enter the input here"))
})();



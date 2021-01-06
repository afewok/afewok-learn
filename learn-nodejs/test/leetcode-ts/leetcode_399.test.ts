function calcEquation(equations: string[][], values: number[], queries: string[][]): number[] {
    let myMap = new Map(),array:number[]=[];
    for (let i = 0; i < values.length; i++) {

        
    }

    for (let i = 0; i < values.length; i++) {
        
    }
    
    return array;
};

/**
 * 399. 除法求值
 */
test('leetcode_399', () => {
    console.log(calcEquation([["a", "b"], ["b", "c"]], [2.0, 3.0], [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]]))
    console.log(calcEquation([["a", "b"], ["b", "c"], ["bc", "cd"]], [1.5, 2.5, 5.0], [["a", "c"], ["c", "b"], ["bc", "cd"], ["cd", "bc"]]))
    console.log(calcEquation([["a", "b"]], [0.5], [["a", "b"], ["b", "a"], ["a", "c"], ["x", "y"]]))
});
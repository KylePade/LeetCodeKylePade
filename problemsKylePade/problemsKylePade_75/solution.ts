/**
 Do not return anything, modify nums in-place instead.
 */
function sortColors(nums: number[]): void {
    
};

export function Solve(inputJsonElement: string): any {
	const inputValues: string[] = inputJsonElement.split("\n");
	const nums: number[] = JSON.parse(inputValues[0]);
	sortColors(nums)
	return nums;
}

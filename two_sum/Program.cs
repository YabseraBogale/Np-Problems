using System.IO.MemoryMappedFiles;
using System.Runtime;

int[] num={3,2,4};
int target=6;

int[] TwoSums(int[] nums,int target)
{
   Dictionary<int,int> map=new Dictionary<int,int>();
   for(int i = 0; i < nums.Length; i++)
    {
        int complement=target-nums[i];
        Console.WriteLine(complement);
        if (map.ContainsKey(complement))
        {
            return new int[] {map[complement],i};
        }
        map[nums[i]]=i;
    } 
    return new int[]{};
}



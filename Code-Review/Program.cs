using System;

namespace Code_Review
{
    class Program
    {
        static void Main(string[] args)
        {
            int size = Convert.ToInt32(Console.ReadLine());
            int[] arr = new int[size];
            for (int i = 0; i < size; i++)
            {
                arr[i] = Convert.ToInt32(Console.ReadLine());
            }

            Sorter.SortArray(arr);
            Console.WriteLine(string.Join(' ', arr));
        }
    }
}

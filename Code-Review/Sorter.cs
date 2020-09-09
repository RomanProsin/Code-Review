namespace Code_Review
{
    public class Sorter
    {
        public static int[] SortArray(int[] array)
        {
            for (int i = 0; i < array.Length - 1; i++)
            {
                for (int j = i + 1; j < array.Length; j++)
                {
                    if (array[i] > array[j])
                    {
                        var buffer = array[j];
                        array[j] = array[i];
                        array[i] = buffer;
                    }
                }
            }

            return array;
        }
    }
}

using Code_Review;
using NUnit.Framework;

namespace Code_Review_Tests
{
    public class Tests
    {
        [TestCase(new int[] { 0, 1, 2 }, new int[] { 0, 1, 2 })]
        [TestCase(new int[] { 0, 0, 4 }, new int[] { 0, 0, 4 })]
        [TestCase(new int[] { 0, 0, -2 }, new int[] { -2, 0, 0 })]
        [TestCase(new int[] { 0, -2, 0 }, new int[] { -2, 0, 0 })]
        [TestCase(new int[] { 0, -2, -2, 1 }, new int[] { -2, -2, 0, 1 })]
        [TestCase(new int[] { 6, 2, 4 }, new int[] { 2, 4, 6 })]

        public static void SortTest(int[] arr, int[] sorted)
        {
            Sorter.SortArray(arr);
            Assert.AreEqual(sorted, arr);
        }
    }
}
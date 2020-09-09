//
//  main.swift
//

import Foundation

var array = readLine()!.components(separatedBy: " ").compactMap{ Int(String($0))! }

for i in 0..<array.count {
    for j in 1..<array.count {
        if array[j] < array[j-1] {
            let tmp = array[j-1]
            array[j-1] = array[j]
            array[j] = tmp
        }
    }
}

print(array)

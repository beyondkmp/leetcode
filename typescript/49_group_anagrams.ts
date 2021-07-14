function groupAnagrams(strs: string[]): string[][] {
  const dicts: Map<string, string[]> = new Map();
  strs.forEach(str => {
    const sortedStr = str.split("").sort().join('');
    if (!dicts.has(sortedStr)) {
      dicts.set(sortedStr, []);
    }
    dicts.set(sortedStr, [...dicts.get(sortedStr), str]);
  })

  return [...dicts.values()]
};

console.log(groupAnagrams([""]));
console.log(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]));
console.log(groupAnagrams(["a"]));


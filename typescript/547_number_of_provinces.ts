class UF {
  private _ids: number[];
  private _count: number;
  constructor(N: number) {
    this._count = N;
    this._ids = new Array<number>(N);
    for (let i = 0; i < N; i++) {
      this._ids[i] = i;
    }
  }

  count(): number {
    return this._count;
  }
  connected(p: number, q: number) {
    return this.find(p) === this.find(q);
  }
  find(p: number): number {
    return this._ids[p];
  }
  union(p: number, q: number) {
    const pID = this.find(p);
    const qID = this.find(q);

    if (pID === qID) {
      return;
    }

    for (let i = 0; i < this._ids.length; i++) {
      if (this._ids[i] === pID) {
        this._ids[i] = qID;
      }
    }
    this._count--;
  }
}
function findCircleNum(isConnected: number[][]): number {
  const uf = new UF(isConnected.length);
  for (let i = 0; i < isConnected.length; i++) {
    for (let j = 0; j < isConnected[i].length; j++) {
      if (isConnected[i][j] === 1) {
        uf.union(i, j);
      }
    }
  }
  return uf.count()
};

const test1: number[][] = [[1,1,0],[1,1,0],[0,0,1]];
console.log(findCircleNum(test1));

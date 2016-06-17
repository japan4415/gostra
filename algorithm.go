package gostra

func Shortest(sg [2]int,ary [][3]int)([]int){
    //変数宣言 level:=重み,ルート flag:=通ったか route:=結果
    level := [][]int{}
    flag := []int{}
    //levelをリセット
    appendAry := []int{0}
    level = append(level,appendAry)
    for i:=0;i<sg[1]-1;i++{
        appendAry = []int{999}
        level = append(level,appendAry)
    }
    //flagをリセット
    for i:=0;i<sg[1];i++{
        flag = append(flag,0)
    }
    current := sg[0]
    addAry := []int{}
    for current<sg[1]{
        //log.Print(current,"について開始します")
        for _,v := range ary{
            if current == v[0]{
                //log.Print(v,"が引っかかりましたa")
                if flag[v[1]-1] == 0{
                    if level[v[1]-1][0]>level[current-1][0]+v[2]{
                        addAry = []int{level[current-1][0] + v[2]}
                        addAry = append(addAry,level[current-1][1:]...)
                        addAry = append(addAry,current)
                        level[v[1]-1] = addAry
                        //log.Print(v[1],"に",level[v[1]-1],"を入れましたa")
                    }
                }
            }
            if current == v[1]{
                //log.Print(v,"が引っかかりましたb")
                if flag[v[0]-1] == 0{
                    if level[v[0]-1][0]>level[current-1][0]+v[2]{
                        addAry = []int{level[current-1][0] + v[2]}
                        addAry = append(addAry,level[current-1][1:]...)
                        addAry = append(addAry,current)
                        level[v[0]-1] = addAry
                        //log.Print(v[0],"に",level[v[0]-1],"を入れましたb")
                    }
                }
            }
            flag[current-1] = 1
        }
        lowest := []int{0,999}
            for kk,vv := range level{
                if lowest[1] > vv[0]{
                    if flag[kk] == 0{
                        lowest[0] = kk + 1
                        lowest[1] = vv[0]
                    }
                }
            }
        //log.Print(lowest)
        current = lowest[0]
        //log.Print(flag,"がフラグ")
        //log.Print(level,"がレベル")
        //log.Print("次は",current,"です")
    }
    //log.Print("出たよ")
    return append(level[sg[1]-1][1:],sg[1])
}
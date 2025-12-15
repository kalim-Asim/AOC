package utils

func Distance(a, b Point) int {
    dx := (a.X - b.X)
    dy := (a.Y - b.Y)
    dz := (a.Z - b.Z)
    return dx*dx + dy*dy + dz*dz
}
package icons


type IconContextType int
const (
    IconContextThreshold IconContextType = 0
    IconContextFixed                     = 1
    IconContextScalable                  = 2
)

type IconContext struct {
    Size      int
    Name      string
    Type      IconContextType
    MaxSize   int               // Type=IconContextScalable
    MinSize   int               // Type=IconContextScalable
    Threshold int               // Type=IconContextThreshold
}

func (self *IconContext) IsValid() bool {
    return (self.Size > 0)
}

def gen_n(n, writer):
    params = dict(n=n, paramlist=', '.join(['T%d' % (i + 1) for i in range(n)]))
    writer.write('// Generated by gen.py. DO NOT modified it\n')
    writer.write('package inject\n')
    writer.write('\n')
    writer.write('import "github.com/gin-gonic/gin"\n')
    writer.write('\n')
    writer.write('func Wrap{n}[{paramlist} any](f func(*gin.Context, {paramlist})) func(*gin.Context) {{\n'.format(**params))
    for i in range(n):
        writer.write('\tgetter{i}, closer{i} := processArg[T{i}]()\n'.format(i=i+1))
    writer.write('\treturn func(c *gin.Context) {\n')
    for i in range(n):
        writer.write('\t\tv{i} := getter{i}(c)\n'.format(i=i+1))
        writer.write('\t\tif closer{i} != nil {{\n'.format(i=i+1))
        writer.write('\t\t\tdefer closer{i}(v{i})\n'.format(i=i+1))
        writer.write('\t\t}\n')
    writer.write('\t\tf(c')
    for i in range(n):
        writer.write(', v{i}'.format(i=i+1))
    writer.write(')\n')
    writer.write('\t}\n')
    writer.write('}\n')

if __name__ == '__main__':
    for n in range(1, 10):
        with open('wrap_{n}.go'.format(n=n), 'w') as f:
            gen_n(n, f)

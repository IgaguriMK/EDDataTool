#! /bin/bash

set -eu

TYPE=$1
echo $TYPE

SRCFILE=out_jcheck/types/$TYPE.json
DESTFILE=model/journal/$TYPE.go

gojson -pkg journal -name $TYPE -input $SRCFILE > $DESTFILE
sed -i $DESTFILE -e 's/Timestamp *string/Timestamp time.Time/'
sed -i $DESTFILE -e '3s/\[\]//'
sed -i $DESTFILE -e '2iimport "time"'
cat >> $DESTFILE <<-EOF
func (e $TYPE) GetEvent() string {
       return e.Event
}

func (e $TYPE) GetTimestamp() time.Time {
       return e.Timestamp
}
EOF


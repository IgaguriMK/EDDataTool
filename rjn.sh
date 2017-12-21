#! /bin/bash

set -eu

if [ $# -lt 1 ]; then
	echo Need arg
	exit 1
fi

ERR=$1

TYPE=`echo $ERR | sed 's/^out_jcheck\/0\.unknown\.\([a-zA-Z]\+\).*$/\1/'`

if [ $TYPE == "" ]; then
	TYPE=VVV
	gvim $ERR &
fi

TMPF=`mktemp`.go
gojson -name $TYPE -subStruct -input $ERR \
	| sed 's/Timestamp *string/Timestamp time.Time/' \
	> $TMPF
cat >> $TMPF <<-EOF

func (e ${TYPE}) GetEvent() string {
	return e.Event
}

func (e ${TYPE}) GetTimestamp() time.Time {
	return e.Timestamp
}

///////////////////////////

func parseWithType(bytes []byte, eventType string) (Event, error) {
	switch eventType {
	////////////////////
	case "${TYPE}":
		var e ${TYPE}
		err := json.Unmarshal(bytes, &e)
		return e, err
	///////////////////
	}
}

EOF
echo 
gvim $TMPF
rm $TMPF

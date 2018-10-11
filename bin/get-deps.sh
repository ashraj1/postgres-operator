#!/bin/bash

# Copyright 2017-2018 Crunchy Data Solutions, Inc.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

echo "Getting project dependencies..."

#sudo yum -y install mercurial golang
which go
if [ $? -eq 1 ]; then
	echo "installing golang..."
	sudo yum -y install golang
fi

which dep
if [ $? -eq 1 ]; then
	echo "installing dep"
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
fi

go get github.com/blang/expenv

#dep ensure

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "pre-pulling container suite images used by the operator..."
$DIR/pre-pull-crunchy-containers.sh
Layers=dao controller service router
SrcFile=user.go activity.go post.go comment.go number.go

MockDir=./internal/test/mock

mock:
	for i in $(Layers); do \
		for j in $(SrcFile); do \
		  echo $$i $$j; \
			mockgen -source=./internal/$$i/$(patsubst %.go,%_$$(i).go,$$j) -destination=$(MockDir)/$(patsubst %.go,%_mock.go,$$(j)) -package=test; \
		done; \
	done


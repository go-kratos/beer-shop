service?=

.PHONY: all api proto wire ent build docker pushi
all: api wire proto ent build
	@ echo "Code generate all done!"

# 生成api
api:
	@if [[ x${service} == x ]]; then \
		find api -maxdepth 3 -mindepth 3 -type d -print | xargs -I {} bash -c 'cd {} && make api'; \
	else \
		find api -maxdepth 3 -mindepth 3 -type d -print | xargs -I {} bash -c  'if [[ {} == *${service}* ]]; then cd {} && make api; fi'; \
	fi

# 生成wire
wire:
	@if [[ x${service} == x ]]; then \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c 'cd {} && make wire'; \
	else \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c  'if [[ {} == *${service}* ]]; then cd {} && make wire; fi'; \
	fi

# 生成config
proto:
	@if [[ x${service} == x ]]; then \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c 'cd {} && make proto'; \
	else \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c  'if [[ {} == *${service}* ]]; then cd {} && make proto; fi'; \
	fi

# 生成ent数据库
ent:
	-@if [[ x${service} == x ]]; then \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c 'cd {} && make ent'; \
	else \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c  'if [[ {} == *${service}* ]]; then cd {} && make ent; fi'; \
	fi

# 编译服务
build:
	@if [[ x${service} == x ]]; then \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c 'cd {} && make build'; \
	else \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c  'if [[ {} == *${service}* ]]; then cd {} && make build; fi'; \
	fi

# 生成docker镜像
docker:
	@if [[ x${service} == x ]]; then \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c 'cd {} && make docker'; \
	else \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c  'if [[ {} == *${service}* ]]; then cd {} && make docker; fi'; \
	fi

# 推送镜像
pushi:
	@if [[ x${service} == x ]]; then \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c 'cd {} && make pushi'; \
	else \
		find app -maxdepth 2 -mindepth 2 -type d -print | xargs -I {} bash -c  'if [[ {} == *${service}* ]]; then cd {} && make pushi; fi'; \
	fi

#!/bin/bash

set -eux

# functions common to more than one distro specific provisioning
url_to_repo() {
    local url="$1"

    local repo=${url#*://}
    repo="${repo//%/}"
    repo="${repo//\//_}"

    echo "$repo"
}

add_repo() {
    local match="$1"
    local add_repo="$2"
    local gpg_check="${3:-true}"

    if [ -z "$match" ]; then
        # we cannot try to add a repo that has no match
        return
    fi

    local repo
    # see if a package we know is in the repo is present
    if repo=$(dnf -y repoquery --qf "%{repoid}" "$1" 2>/dev/null | grep ..\*); then
        DNF_REPO_ARGS+=" --enablerepo=$repo"
    else
        local repo_url="${REPOSITORY_URL}${add_repo}"
        local repo_name
        repo_name=$(url_to_repo "$repo_url")
        if ! dnf repolist | grep "$repo_name"; then
            dnf config-manager --add-repo="${repo_url}" >&2
            if ! $gpg_check; then
                disable_gpg_check "$add_repo" >&2
            fi
        fi
        DNF_REPO_ARGS+=" --enablerepo=$repo_name"
    fi
}

add_group_repo() {
    local match="$1"

    add_repo "$match" "$DAOS_STACK_GROUP_REPO"
    group_repo_post
}

add_local_repo() {
    add_repo 'argobots' "$DAOS_STACK_LOCAL_REPO" false
}

disable_gpg_check() {
    local url="$1"

    repo="$(url_to_repo "$url")"
    # bug in EL7 DNF: this needs to be enabled before it can be disabled
    dnf config-manager --save --setopt="$repo".gpgcheck=1
    dnf config-manager --save --setopt="$repo".gpgcheck=0
    # but even that seems to be not enough, so just brute-force it
    if [ -d /etc/yum.repos.d ] &&
       ! grep gpgcheck /etc/yum.repos.d/"$repo".repo; then
        echo "gpgcheck=0" >> /etc/yum.repos.d/"$repo".repo
    fi
}

dump_repos() {
        for file in "$REPOS_DIR"/*.repo; do
            echo "---- $file ----"
            cat "$file"
        done
}

env > /root/last_run-env.txt
if ! grep ":$MY_UID:" /etc/group; then
  groupadd -g "$MY_UID" jenkins
fi
mkdir -p /localhome
if ! grep ":$MY_UID:$MY_UID:" /etc/passwd; then
  useradd -b /localhome -g "$MY_UID" -u "$MY_UID" -s /bin/bash jenkins
fi
jenkins_ssh=/localhome/jenkins/.ssh
mkdir -p "${jenkins_ssh}"
if ! grep -q -s -f /tmp/ci_key.pub "${jenkins_ssh}/authorized_keys"; then
  cat /tmp/ci_key.pub >> "${jenkins_ssh}/authorized_keys"
fi
root_ssh=/root/.ssh
if ! grep -q -f /tmp/ci_key.pub "${root_ssh}/authorized_keys"; then
  cat /tmp/ci_key.pub >> "${root_ssh}/authorized_keys"
fi
cp /tmp/ci_key.pub "${jenkins_ssh}/id_rsa.pub"
cp /tmp/ci_key "${jenkins_ssh}/id_rsa"
cp /tmp/ci_key_ssh_config "${jenkins_ssh}/config"
chmod 700 "${jenkins_ssh}"
chmod 600 "${jenkins_ssh}"/{authorized_keys,id_rsa*,config}
chown -R jenkins.jenkins /localhome/jenkins/
echo "jenkins ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/jenkins

# remove any defined releasever
# any deviation from the default releasever for a distro should be
# handled on a per-test-run, or even per-dnf command basis
rm -f  /etc/yum/vars/releasever

# defined in ci/functional/post_provision_config_nodes_<distro>.sh
# and catted to the remote node along with this script
post_provision_config_nodes

systemctl enable nfs-server.service
systemctl start nfs-server.service
sync
sync
exit 0

#!/bin/bash

export TF_VAR_billing_account=${USER}
export TF_ADMIN=${USER}-terraform-admin
export TF_CREDS=~/.config/gcloud/${USER}-terraform-admin.json
